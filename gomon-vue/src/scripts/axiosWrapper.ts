import axios from 'axios'

const ServerInfo = {
  baseURL: '127.0.0.1:8080',
  timeout: 4000, // ms
}

const AxiosObject = axios.create(ServerInfo)

export default AxiosObject
