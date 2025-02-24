import axios from 'axios';
import config from './config';

const instance = axios.create({
  baseURL: config.server.api_url,
  timeout: config.api.timeout_ms,
  headers: {
    'Content-Type': 'application/json',
  },
});

export default instance; 