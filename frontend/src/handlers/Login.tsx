import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

const Login: React.FC = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setErrorMessage('');
        try {
            const response = await axios.post(
                'http://localhost:8000/login',
                { username, password },
                { withCredentials: true }
            );
            console.log('ログイン成功:', response.data);
            alert('ログインが完了しました。');
            // ログイン成功時の処理（例：リダイレクト）を追加

            localStorage.setItem('authToken', response.data.token);
            navigate('/home');
        } catch (error) {
            if (axios.isAxiosError(error)) {
                console.error('ログインエラー:', error.response?.data || error.message);
                setErrorMessage(error.response?.data?.message || 'ログインに失敗しました。');
            } else {
                console.error('予期せぬエラー:', error);
                setErrorMessage('予期せぬエラーが発生しました。');
            }
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