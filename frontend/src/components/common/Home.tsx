import React, { useEffect } from 'react';
import { Navigate } from 'react-router-dom';
import Sidebar from './Sidebar';
import { ContentContainer } from '../../styles/SidebarStyles';
import { log } from '../Log/Logging';

const isAuthenticated = (): boolean => {
    return localStorage.getItem('authToken') !== null;
};

// みんなの制作物が表示されるホームページ
// いいね機能があるといいかも

const Home: React.FC = () => {
    // const userId = localStorage.getItem('userId');
    useEffect(() => {
        log('Homeコンポーネントがマウントされました', 'info');
    }, []);

    // ユーザーが認証されていない場合はログインページにリダイレクト
    if (!isAuthenticated()) {
        return <Navigate to="/Signin" replace />;
    }

    return (
        <div>
            <Sidebar />
            <ContentContainer>
                <h1>Home Page</h1>
                <p>Welcome to the Home Page!</p>
                <p>追加予定</p>
                {/* {userId ? (
                    <Link to={`/${userId}`}>プロフィールはこちら</Link>
                ) : (
                    <p>ユーザー情報が見つかりません。</p>
                )} */}
            </ContentContainer>
        </div>
    );
};

export default Home;
