export const log = (message: string, level: 'debug' | 'info' | 'warn' | 'error' = 'info') => {
    console.log(`[${level.toUpperCase()}] ${message}`);
  };
  