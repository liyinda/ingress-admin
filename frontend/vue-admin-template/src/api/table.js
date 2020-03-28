import request from '@/utils/request'

export function getList(token) {
  return request({
    url: '/home/table',
    method: 'get',
    params: { token }
  })
}

export function getAnnotations(token) {
  return request({
    url: '/home/annotations',
    method: 'get',
    params: { token }
  })
}

export function addDomain(params) {
  return request({
    url: '/home/add',
    method: 'post',
    params: params
  })
}

export function updateDomain(params) {
  return request({
    url: '/home/update',
    method: 'post',
    params: params
  })
}

export function delDomain(params) {
  return request({
    url: '/home/del',
    method: 'post',
    params: params
  })
}
