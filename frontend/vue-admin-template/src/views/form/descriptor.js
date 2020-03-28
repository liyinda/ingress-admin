export default {
  object: {
    type: 'object',
    label: 'Ingress Object',
    fields: {
      name: { type: 'string', required: true, hidden: false },
      namespace: { type: 'string', required: true, hidden: false },
      host: { type: 'string', required: true, hidden: false }
    }
  },
  hashmap: {
    type: 'object',
    label: 'annotations',
    defaultField: {
      type: 'string',
      required: true
    }
  },
  children: {
    type: 'array',
    label: 'Ingress Paths',
    defaultField: {
      type: 'object',
      fields: {
        path: { type: 'string', required: true },
        servicename: { type: 'string', required: true },
        serviceport: { type: 'string', required: true }
      }
    }
  }
}
