// components/Logout.tsx

import React, { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

const Logout: React.FC = () => {
    const navigate = useNavigate();

    useEffect(() => {
        localStorage.removeItem('authToken');
        navigate('/login');
    }, [navigate]);

    return (
        <div>
            <p>ログアウト中...</p>
        </div>
    );
};

export default Logout;
