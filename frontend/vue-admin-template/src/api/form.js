import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/home/table',
    method: 'get',
    params
  })
}
