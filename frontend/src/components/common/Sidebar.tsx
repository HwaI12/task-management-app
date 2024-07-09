import React from 'react';
import {
    SidebarContainer,
    SidebarHeader,
    SidebarMenu,
    MenuItem,
    SidebarLogoImg,
    StyledIoHomeOutline,
    StyledIoNotificationsOutline,
    StyledIoPersonOutline,
    StyledIoLogOutOutline,
    StyledLink,
    StyledIoAddCircleOutline
} from '../../styles/SidebarStyles';
import logo from '../../assets/T.png';

const Sidebar: React.FC = () => {
    const userId = localStorage.getItem('userId');

    return (
        <SidebarContainer>
            <SidebarHeader>
                <SidebarLogoImg src={logo} alt="logo" />
            </SidebarHeader>
            <SidebarMenu>
                <MenuItem>
                    <StyledLink to="/home"><StyledIoHomeOutline /></StyledLink>
                </MenuItem>
                <MenuItem>
                    <StyledLink to="/Notifications"><StyledIoNotificationsOutline /></StyledLink>
                </MenuItem>
                <MenuItem>
                    {userId && (
                        <StyledLink to={`/${userId}/addtask`}><StyledIoAddCircleOutline /></StyledLink>
                    )}
                </MenuItem>
                <MenuItem>
                    {userId && (
                        <StyledLink to={`/${userId}`}><StyledIoPersonOutline /></StyledLink>
                    )}
                </MenuItem>
                <MenuItem>
                    <StyledLink to="/logout"><StyledIoLogOutOutline /></StyledLink>
                </MenuItem>
            </SidebarMenu>
        </SidebarContainer>
    );
};

export default Sidebar;
