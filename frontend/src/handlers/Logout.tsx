// components/Logout.tsx

import React, { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

const Logout: React.FC = () => {
    const navigate = useNavigate();

    useEffect(() => {
        // localStorageから認証トークンを削除する
        localStorage.removeItem('authToken');

        // ログインページにリダイレクトする
        navigate('/login');
    }, [navigate]);

    return (
        <div>
            <p>ログアウト中...</p>
            {/* 任意でここにローディングスピナーやメッセージを追加できます */}
        </div>
    );
};

export default Logout;
