import yaml from 'js-yaml'
import fs from 'fs'

interface Config {
  server: {
    api_url: string;
    frontend_url: string;
  };
  auth: {
    token_key: string;
  };
  api: {
    timeout_ms: number;
    retry_attempts: number;
  };
}

const config: Config = yaml.load(
  fs.readFileSync('./config/config.yaml', 'utf8')
) as Config;

export default config; 