import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { Container, Form, Title, Label, Input, Button, LinkText } from '../styles/UserAuthStyles';
import { Link } from 'react-router-dom';

const isAuthenticated = (): boolean => {
    return localStorage.getItem('authToken') !== null;
};

const Signup: React.FC = () => {
    const [user_id, setUserId] = useState('');
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();

    useEffect(() => {
        if (isAuthenticated()) {
            navigate('/home');
        }
    }, [navigate]);

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setErrorMessage('');
        try {
            const response = await axios.post('http://localhost:8000/register', {
                user_id,
                username,
                email,
                password_hash: password,
            });

            console.log('Resister successful:', response.data);
            localStorage.setItem('authToken', response.data.token);
            navigate('/home');
        } catch (error) {
            if (axios.isAxiosError(error)) {
                console.error('Signup error:', error.response?.data || error.message);
                setErrorMessage(error.response?.data?.message || '登録に失敗しました。');
            } else {
                console.error('Unexpected error:', error);
                setErrorMessage('予期せぬエラーが発生しました。');
            }
        }
    }

    return (
        <Container>
            <Form>
                <Title>新規登録</Title>
                {errorMessage && <div style={{ color: 'red' }}>{errorMessage}</div>}
                <form onSubmit={handleSubmit}>
                    <Label htmlFor="user_id">ユーザーID</Label>
                    <Input
                        id="user_id"
                        type="text"
                        placeholder="tnk123"
                        value={user_id}
                        onChange={(e) => setUserId(e.target.value)}
                        required
                    />
                    <Label htmlFor="username">ユーザー名</Label>
                    <Input
                        id="username"
                        type="text"
                        placeholder="tanaka"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                        required
                    />
                    <Label htmlFor="password">メールアドレス</Label>
                    <Input
                        id="email"
                        type="email"
                        placeholder="example@gmail.com"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                    />
                    <Label htmlFor="password">パスワード</Label>
                    <Input
                        id="password"
                        type="password"
                        placeholder="password123"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                    <Button type="submit">登録</Button>
                </form>
                <LinkText>
                    <Link to="/Signin">ログインはこちら</Link>
                </LinkText>
            </Form>
        </Container>
    );
};

export default Signup;
