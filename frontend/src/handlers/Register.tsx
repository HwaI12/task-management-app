// components/Register.jsx

import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

const Register = () => {
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        try {
            const response = await axios.post('http://localhost:8000/register', {
                username,
                email,
                password_hash: password,
            });

            if (response.status === 201) {
                console.log('登録成功');
                alert('ユーザー登録が完了しました。');

                // Assuming you receive an authentication token in the response
                // Set the token in localStorage (you need to implement this part)
                localStorage.setItem('authToken', response.data.token);
                navigate('/home');
            } else {
                console.error('登録失敗:', response.status);
                setErrorMessage('ユーザーの登録に失敗しました。');
            }
        } catch (error) {
            console.error('ネットワークエラー:', error);
            setErrorMessage('ネットワークエラーが発生しました。');
        }
    };

    return (
        <div>
            <h2>新規登録</h2>
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
                    type="email"
                    placeholder="Email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                />
                <input
                    type="password"
                    placeholder="パスワード"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                />
                <button type="submit">登録</button>
            </form>
        </div>
    );
};

export default Register;
