import React from 'react';
import { Navigate, Link } from 'react-router-dom';
import Sidebar from './Sidebar';
import { ContentContainer } from '../../styles/SidebarStyles';

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
            <Sidebar />
            <ContentContainer>
                <h2>Home Page</h2>
                <p>Welcome to the Home Page!</p>
                {userId ? (
                    <Link to={`/${userId}`}>プロフィールはこちら</Link>
                ) : (
                    <p>ユーザー情報が見つかりません。</p>
                )}
            </ContentContainer>
        </div>
    );
};

export default Home;
