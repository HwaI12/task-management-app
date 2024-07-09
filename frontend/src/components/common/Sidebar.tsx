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
    StyledIoLogOutOutline
} from '../../styles/SidebarStyles';
import logo from '../../assets/T.png';

const Sidebar: React.FC = () => {
    return (
        <SidebarContainer>
            <SidebarHeader>
                <SidebarLogoImg src={logo} alt="logo" />
            </SidebarHeader>
            <SidebarMenu>
                <MenuItem>
                    <StyledIoHomeOutline />
                </MenuItem>
                <MenuItem>
                    <StyledIoNotificationsOutline />
                </MenuItem>
                <MenuItem>
                    <StyledIoPersonOutline />
                </MenuItem>
                <MenuItem>
                    <StyledIoLogOutOutline />
                </MenuItem>
            </SidebarMenu>
        </SidebarContainer>
    );
};

export default Sidebar;
