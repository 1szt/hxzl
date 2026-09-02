// ============================================================
//  SBCC 控制中心 — 前端应用
//  功能：登录/注册/注销/状态查询 + WebSocket 实时数据
// ============================================================

// ─── DOM 引用 ───────────────────────────────────────────
const loginOverlay = document.getElementById('loginOverlay')
const mainContent  = document.getElementById('mainContent')
const loginForm    = document.getElementById('loginForm')
const registerForm = document.getElementById('registerForm')
const loginError   = document.getElementById('loginError')
const registerError = document.getElementById('registerError')
const userName     = document.getElementById('userName')
const btnLogout    = document.getElementById('btnLogout')
const tabs         = document.querySelectorAll('.login-tab')
const dataValue    = document.getElementById('dataValue')
const dataTime     = document.getElementById('dataTime')
const dataCard     = document.getElementById('dataCard')
const historyList  = document.getElementById('historyList')
const statusDot    = document.getElementById('statusDot')
const statusText   = document.getElementById('statusText')

// ─── 工具函数 ──────────────────────────────────────────

function apiURL(path) {
  // 自动使用当前页面的 host（端口跟随服务器配置）
  return window.location.origin + '/api/user' + path
}

function showError(el, msg) {
  el.textContent = msg
}

function clearError(el) {
  el.textContent = ''
}

// ─── Tab 切换 ──────────────────────────────────────────

tabs.forEach(tab => {
  tab.addEventListener('click', () => {
    tabs.forEach(t => t.classList.remove('active'))
    tab.classList.add('active')

    const target = tab.dataset.tab
    if (target === 'login') {
      loginForm.style.display = ''
      registerForm.style.display = 'none'
      clearError(loginError)
    } else {
      loginForm.style.display = 'none'
      registerForm.style.display = ''
      clearError(registerError)
    }
  })
})

// ─── 登录 ──────────────────────────────────────────────

loginForm.addEventListener('submit', async (e) => {
  e.preventDefault()
  clearError(loginError)

  const username = document.getElementById('loginUsername').value.trim()
  const password = document.getElementById('loginPassword').value

  try {
    const res = await fetch(apiURL(''), {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, password }),
    })

    const data = await res.json()

    if (data.ok) {
      await checkLoginStatus()
    } else {
      showError(loginError, data.message || '登录失败')
    }
  } catch (err) {
    showError(loginError, '网络错误，请重试')
  }
})

// ─── 注册 ──────────────────────────────────────────────

registerForm.addEventListener('submit', async (e) => {
  e.preventDefault()
  clearError(registerError)

  const username = document.getElementById('registerUsername').value.trim()
  const password = document.getElementById('registerPassword').value

  if (password.length < 6) {
    showError(registerError, '密码长度不能少于6位')
    return
  }

  try {
    const res = await fetch(apiURL('/register'), {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, password }),
    })

    const data = await res.json()

    if (data.ok) {
      await checkLoginStatus()
    } else {
      showError(registerError, data.message || '注册失败')
    }
  } catch (err) {
    showError(registerError, '网络错误，请重试')
  }
})

// ─── 注销 ──────────────────────────────────────────────

btnLogout.addEventListener('click', async () => {
  try {
    const res = await fetch(apiURL('/logout'), {
      method: 'POST',
    })

    const data = await res.json()

    if (data.ok) {
      // 断开 WebSocket
      if (window.ws) {
        window.ws.close()
        window.ws = null
      }
      showLoginUI()
    }
  } catch (err) {
    console.error('注销失败:', err)
  }
})

// ─── 发送测试消息 ────────────────────────────────────

document.getElementById('btnTestMsg').addEventListener('click', () => {
  if (window.ws && window.ws.readyState === WebSocket.OPEN) {
    window.ws.send(JSON.stringify({hello: 'world', time: new Date().toLocaleTimeString()}))
  }
})

// ─── 检查登录状态 ──────────────────────────────────────

async function checkLoginStatus() {
  try {
    const res = await fetch(apiURL('/me'))
    const data = await res.json()

    if (data.logged_in) {
      showMainUI(data.username)
    } else {
      showLoginUI()
    }
  } catch (err) {
    console.error('检查登录状态失败:', err)
    showLoginUI()
  }
}

// ─── UI 切换 ───────────────────────────────────────────

function showLoginUI() {
  loginOverlay.style.display = ''
  mainContent.style.display  = 'none'

  // 清空表单
  loginForm.reset()
  registerForm.reset()
  clearError(loginError)
  clearError(registerError)

  // 重置到登录 Tab
  tabs.forEach(t => t.classList.remove('active'))
  document.querySelector('.login-tab[data-tab="login"]').classList.add('active')
  loginForm.style.display = ''
  registerForm.style.display = 'none'
}

function showMainUI(username) {
  loginOverlay.style.display = 'none'
  mainContent.style.display  = ''
  userName.textContent = username || '用户'

  // 连接 WebSocket
  connectWebSocket()
}

// ─── WebSocket ─────────────────────────────────────────

function connectWebSocket() {
  // 如果已有连接，先关闭
  if (window.ws) {
    try { window.ws.close() } catch (_) {}
    window.ws = null
  }

  const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'
  const wsURL = protocol + '//' + window.location.host + '/ws'

  const ws = new WebSocket(wsURL)
  window.ws = ws

  ws.onopen = () => {
    statusDot.className = 'status-dot connected'
    statusText.textContent = '已连接'
  }

  ws.onclose = () => {
    statusDot.className = 'status-dot'
    statusText.textContent = '已断开，5秒后重连...'
    window.ws = null
    // 自动重连
    setTimeout(connectWebSocket, 5000)
  }

  ws.onerror = () => {
    statusDot.className = 'status-dot'
    statusText.textContent = '连接异常'
  }

  ws.onmessage = (event) => {
    try {
      const msg = JSON.parse(event.data)
      handleWSMessage(msg)
    } catch (err) {
      console.error('WS 消息解析失败:', err)
    }
  }
}

// ─── 处理 WS 消息 ──────────────────────────────────────

const MAX_HISTORY = 20

function handleWSMessage(msg) {
  switch (msg.type) {
    case 'home.random':
      handleHomeRandom(msg.payload)
      break
    // 后续模块可以在这里添加更多 case
    default:
      console.log('未知消息类型:', msg.type, msg.payload)
  }
}

function handleHomeRandom(payload) {
  const value = payload.value
  const time  = payload.time || ''

  // 更新当前数值
  dataValue.textContent = value
  dataTime.textContent  = time ? '更新于 ' + time : ''

  // 动画效果
  dataValue.classList.remove('pop')
  void dataValue.offsetWidth // 触发回流以重播动画
  dataValue.classList.add('pop')
  dataCard.classList.add('has-data')

  // 添加到历史记录
  const emptyMsg = historyList.querySelector('.empty-msg')
  if (emptyMsg) emptyMsg.remove()

  const item = document.createElement('div')
  item.className = 'history-item'
  item.innerHTML = `
    <span class="h-time">${time}</span>
    <span class="h-val">${value}</span>
  `
  historyList.prepend(item)

  // 限制历史记录条数
  while (historyList.children.length > MAX_HISTORY) {
    historyList.removeChild(historyList.lastChild)
  }
}

// ─── 启动 ──────────────────────────────────────────────

document.addEventListener('DOMContentLoaded', () => {
  checkLoginStatus()
})
