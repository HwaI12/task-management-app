import React from 'react';
import { Navigate } from 'react-router-dom';
import { Link } from 'react-router-dom';

const isAuthenticated = (): boolean => {
    return localStorage.getItem('authToken') !== null;
};

const Home: React.FC = () => {
    const userId = localStorage.getItem('userId');

    // ユーザーが認証されていない場合はログインページにリダイレクト
    if (!isAuthenticated()) {
        return <Navigate to="/Signin" replace />;
    }

    return (
        <div>
            <h2>Home Page</h2>
            <p>Welcome to the Home Page!</p>
            {userId ? (
                <Link to={`/${userId}`}>プロフィールはこちら</Link>
            ) : (
                <p>ユーザー情報が見つかりません。</p>
            )}
        </div>
    );
};

export default Home;
