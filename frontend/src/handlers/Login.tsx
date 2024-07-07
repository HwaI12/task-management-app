import React, { useState } from 'react';
import axios from 'axios';

const Login: React.FC = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [errorMessage, setErrorMessage] = useState('');

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        try {
            const response = await axios.post(
                'http://localhost:8000/login',
                {
                    username,
                    password,
                },
                { withCredentials: true } // クッキーを送信するために重要
            );

            console.log('Response:', response);

            if (response.status === 200) {
                console.log('ログイン成功');
                alert('ログインが成功しました');
            } else {
                console.error('ログイン失敗:', response.status);
                setErrorMessage('ユーザー名またはパスワードが正しくありません');
            }
        } catch (error) {
            console.error('ネットワークエラー:', error);
            setErrorMessage('ネットワークエラーが発生しました');
        }
    };

    return (
        <div>
            <h2>ログイン</h2>
            {errorMessage && <div style={{ color: 'red' }}>{errorMessage}</div>}
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    placeholder="Username"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    required
                />
                <input
                    type="password"
                    placeholder="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                />
                <button type="submit">Login</button>
            </form>
        </div>
    );
};

export default Login;
