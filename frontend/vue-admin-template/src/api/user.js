import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/passport/login',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/home/userinfo',
    method: 'get',
    params: { token }
  })
}

export function logout() {
  return request({
    url: '/passport/logout',
    method: 'post'
  })
}
