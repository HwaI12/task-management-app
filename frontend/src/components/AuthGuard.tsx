// AuthGuard.tsx

import React, { ReactNode } from 'react';
import { Navigate } from 'react-router-dom';

const isAuthenticated = (): boolean => {
    return localStorage.getItem('authToken') !== null;
};

interface AuthGuardProps {
    children: ReactNode;
}

const AuthGuard: React.FC<AuthGuardProps> = ({ children }) => {
    if (!isAuthenticated()) {
        return <Navigate to="/login" replace />;
    }

    return <>{children}</>;
};

export default AuthGuard;
