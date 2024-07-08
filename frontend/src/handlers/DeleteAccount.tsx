import React, { useState, useEffect } from 'react';
import { Navigate } from 'react-router-dom';

const isAuthenticated = (): boolean => {
    return localStorage.getItem('authToken') !== null;
};

const DeleteAccount: React.FC = () => {
    const [userId, setUserId] = useState<string>('');
    const [password, setPassword] = useState<string>('');
    const [errorMessage, setErrorMessage] = useState<string>('');

    useEffect(() => {
        // ユーザーが認証されているかどうかの確認
        const checkAuthentication = async () => {
            try {
                // 認証状態の確認を行うためのAPIリクエストなどを実行
                const response = await fetch('http://localhost:8000/check-auth', {
                    method: 'GET',
                    credentials: 'include',
                });
                if (response.ok) {
                    // 認証された場合、ユーザーIDを取得する処理を実行
                    fetchUserId();
                } else {
                    // 認証されていない場合はログインページにリダイレクト
                    console.error('Authentication check failed');
                    setErrorMessage('認証されていません。');
                }
            } catch (error) {
                console.error('Network error:', error);
                setErrorMessage('ネットワークエラーが発生しました。');
            }
        };

        checkAuthentication();
    }, []);

    const fetchUserId = async () => {
        try {
            const response = await fetch('http://localhost:8000/user_id', {
                method: 'GET',
                credentials: 'include',
            });
            if (response.ok) {
                const data = await response.json();
                setUserId(data.user_id);
            } else {
                console.error('Failed to fetch user ID');
                setErrorMessage('ユーザーIDの取得に失敗しました。');
            }
        } catch (error) {
            console.error('Network error:', error);
            setErrorMessage('ネットワークエラーが発生しました。');
        }
    };

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setErrorMessage('');

        // ユーザーIDとパスワードをサーバーに送信して削除処理を行う
        try {
            const response = await fetch('http://localhost:8000/delete', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ user_id: userId, password }),
                credentials: 'include',
            });

            if (response.ok) {
                console.log('Account deletion successful');
                alert('アカウントが正常に削除されました。');
            } else {
                const errorData = await response.json();
                console.error('Account deletion failed:', errorData.message);
                setErrorMessage(errorData.message || 'アカウントの削除に失敗しました。');
            }
        } catch (error) {
            console.error('Network error:', error);
            setErrorMessage('ネットワークエラーが発生しました。');
        }
    };

    if (!isAuthenticated()) {
        // If not authenticated, redirect to /login
        return <Navigate to="/login" replace />;
    }

    return (
        <div>
            <h2>退会</h2>
            {errorMessage && <div style={{ color: 'red' }}>{errorMessage}</div>}
            {userId && (
                <div>
                    <p>以下のユーザーIDで退会します。間違いがないことを確認してください。</p>
                    <p>User ID: {userId}</p>
                </div>
            )}
            <form onSubmit={handleSubmit}>
                <input
                    type="password"
                    placeholder="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                />
                <button type="submit">退会</button>
            </form>
        </div>
    );
};

export default DeleteAccount;
