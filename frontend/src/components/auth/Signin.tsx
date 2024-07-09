import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { Container, Form, Title, Label, Input, Button, LinkText } from '../../styles/UserAuthStyles';
import { Link } from 'react-router-dom';

const isAuthenticated = (): boolean => {
    return localStorage.getItem('authToken') !== null;
};

const Signin: React.FC = () => {
    const [user_id, setUser_id] = useState('');
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
            const response = await axios.post(
                'http://localhost:8000/login',
                { user_id, password },
                { withCredentials: true }
            );
            console.log('Signin successful:', response.data);
            localStorage.setItem('authToken', response.data.token);
            localStorage.setItem('userId', user_id);
            localStorage.setItem('username', response.data.username);
            navigate('/home');
        } catch (error) {
            if (axios.isAxiosError(error)) {
                console.error('Signin error:', error.response?.data || error.message);
                setErrorMessage(error.response?.data?.message || 'ログインに失敗しました。');
            } else {
                console.error('Unexpected error:', error);
                setErrorMessage('予期せぬエラーが発生しました。');
            }
        }
    };

    return (
        <Container>
            <Form>
                <Title>ログイン</Title>
                {errorMessage && <div style={{ color: 'red', marginBottom: '1rem' }}>{errorMessage}</div>}
                <form onSubmit={handleSubmit}>
                    <Label htmlFor="user_id">ユーザーID</Label>
                    <Input
                        id="user_id"
                        type="text"
                        placeholder="ユーザーIDを入力してください"
                        value={user_id}
                        onChange={(e) => setUser_id(e.target.value)}
                        required
                    />
                    <Label htmlFor="password">パスワード</Label>
                    <Input
                        id="password"
                        type="password"
                        placeholder="パスワードを入力してください"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                    <Button type="submit">ログイン</Button>
                </form>
                <LinkText>
                    <Link to="/Signup">会員登録はこちら</Link>
                </LinkText>
            </Form>
        </Container>
    );
};

export default Signin;
