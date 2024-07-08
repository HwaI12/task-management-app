// components/Logout.tsx

import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import styled from 'styled-components';

const LogoutContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background-color: #fff;
  padding: 20px;
  transition: background-color 0.3s ease-in-out;
  box-sizing: border-box;
`;

const ConfirmationBox = styled.div`
  background: white;
  padding: 2rem;
  border-radius: 10px;
  width: 100%;
  max-width: 400px;
  box-sizing: border-box;
  margin: auto;
`;

const Button = styled.button`
  width: 100%;
  height: 3rem;
  padding: 0.75rem;
  margin-bottom: 1rem;
  border: none;
  border-radius: 5px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s ease-in-out, transform 0.2s ease-in-out;
  `;

const ConfirmButton = styled(Button)`
  background-color: #eede77;
  color: #333;

  &:hover {
    background-color: #EBD961;
  }
`;

const CancelButton = styled(Button)`
  background-color: #f0f0f0;
  color: #333;

  &:hover {
    background-color: #E8E8E8;
  }
`;

const Logout: React.FC = () => {
    const navigate = useNavigate();
    const [isConfirming, setIsConfirming] = useState(true);

    const handleLogout = () => {
        localStorage.removeItem('authToken');
        setIsConfirming(false);
        setTimeout(() => navigate('/Signin'), 2000);
    };

    const handleCancel = () => {
        navigate('/home');
    };

    if (!isConfirming) {
        return (
            <LogoutContainer>
                <ConfirmationBox>
                    <p>ログアウトしました。<br></br>サインインページに戻ります...</p>
                </ConfirmationBox>
            </LogoutContainer>
        );
    }

    return (
        <LogoutContainer>
            <ConfirmationBox>
                <h2>ログアウトの確認</h2>
                <p>本当にログアウトしますか？</p>
                <ConfirmButton onClick={handleLogout}>はい</ConfirmButton>
                <CancelButton onClick={handleCancel}>いいえ</CancelButton>
            </ConfirmationBox>
        </LogoutContainer>
    );
};

export default Logout;