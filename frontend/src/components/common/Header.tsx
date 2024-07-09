import React from 'react';
import { useNavigate } from 'react-router-dom';
import { HeaderContainer, Logo, Button } from '../../styles/HeaderStyles';

const Header: React.FC = () => {
    const navigate = useNavigate();
  
    const handleLogout = () => {
      navigate('/logout');
    };

    const handleHome = () =>{
        navigate('/Home');
      };
  
    return (
      <HeaderContainer>
        <Logo onClick={handleHome}>TaskPal</Logo>
        <Button onClick={handleLogout}>ログアウト</Button>
      </HeaderContainer>
    );
  };
  
  export default Header;