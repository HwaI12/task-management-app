// AuthGuard.tsx

import React, { ReactNode } from 'react';
import { Navigate } from 'react-router-dom';

const isAuthenticated = (): boolean => {
    return localStorage.getItem('authToken') !== null;
};

interface AuthGuardProps {
    children: ReactNode;
}

// 認証されていない場合はログインページにリダイレクトするコンポーネント
const AuthGuard: React.FC<AuthGuardProps> = ({ children }) => {
    if (!isAuthenticated()) {
        return <Navigate to="/Signin" replace />;
    }

    return <>{children}</>;
};

export default AuthGuard;
