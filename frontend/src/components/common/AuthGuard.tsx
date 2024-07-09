import React, { ReactNode } from 'react';
import { Navigate, useParams } from 'react-router-dom';

// ユーザーが認証されているか確認する関数
const isAuthenticated = (): boolean => {
    return localStorage.getItem('authToken') !== null;
};

// 認証されたユーザーIDを取得する関数
const getAuthenticatedUserId = (): string | null => {
    return localStorage.getItem('userId');
};

interface AuthGuardProps {
    children: ReactNode;
    requireMatch?: boolean;
}

// 認証されていない場合はログインページにリダイレクトするコンポーネント
const AuthGuard: React.FC<AuthGuardProps> = ({ children, requireMatch }) => {
    const { userId } = useParams<{ userId: string }>();

    if (!isAuthenticated()) {
        return <Navigate to="/Signin" replace />;
    }

    if (requireMatch && userId !== getAuthenticatedUserId()) {
        return <Navigate to="/Home" replace />;
    }

    return <>{children}</>;
};

export default AuthGuard;
