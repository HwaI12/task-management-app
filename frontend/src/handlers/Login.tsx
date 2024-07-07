import React, { useState } from 'react';
import axios from 'axios';

const Login: React.FC = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [errorMessage, setErrorMessage] = useState('');

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8000/login', {
        username,
        password,
      });

      if (response.status === 200) {
        console.log('ログイン成功');
        // ログイン成功時の処理を追加
        alert('ログインが成功しました。');
      } else {
        console.error('ログイン失敗:', response.status);
        // ログイン失敗時の処理を追加
        setErrorMessage('ユーザー名またはパスワードが正しくありません。');
      }
    } catch (error) {
      console.error('ネットワークエラー:', error);
      // ネットワークエラー時の処理を追加
      setErrorMessage('ネットワークエラーが発生しました。');
    }
  };

  return (
    <div>
      <h2>ログイン</h2>
      {errorMessage && <div style={{ color: 'red' }}>{errorMessage}</div>}
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="ユーザー名"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          required
        />
        <input
          type="password"
          placeholder="パスワード"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <button type="submit">ログイン</button>
      </form>
    </div>
  );
};

export default Login;
