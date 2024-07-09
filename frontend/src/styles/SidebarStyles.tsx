import styled from 'styled-components';
import {
    IoHomeOutline,
    IoNotificationsOutline,
    IoPersonOutline,
    IoLogOutOutline
} from "react-icons/io5";

export const SidebarContainer = styled.div`
    width: 60px;
    background-color: #f8f8f8;
    height: 100vh;
    box-shadow: 2px 0 5px rgba(0,0,0,0.1);
    display: flex;
    flex-direction: column;
    position: fixed;
    top: 0;
    left: 0;
    z-index: 1000;
`;

export const SidebarHeader = styled.div`
    padding: 20px;
    background-color: #fff;
    box-shadow: 0 2px 5px rgba(0,0,0,0.1);
    display: flex;
    justify-content: center;
    align-items: center;
`;

export const SidebarMenu = styled.div`
    flex: 1;
    padding: 20px 0;
    overflow-y: auto;
    display: flex;
    flex-direction: column;
    align-items: center;
`;

export const MenuItem = styled.div`
    width: 100%;
    margin-bottom: 20px;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 10px;
    cursor: pointer;
    transition: background-color 0.3s;

    &:hover {
        background-color: #eee;
    }
`;

export const SidebarLogoImg = styled.img`
    width: 50px;
    height: auto; /* 高さを自動調整 */
    object-fit: cover; /* アスペクト比を維持 */
`;

export const StyledIoHomeOutline = styled(IoHomeOutline)`
    font-size: 24px;
    color: #333;
`;

export const StyledIoNotificationsOutline = styled(IoNotificationsOutline)`
    font-size: 24px;
    color: #333;
`;

export const StyledIoPersonOutline = styled(IoPersonOutline)`
    font-size: 24px;
    color: #333;
`;

export const StyledIoLogOutOutline = styled(IoLogOutOutline)`
    font-size: 24px;
    color: #333;
`;

export const ContentContainer = styled.div`
    margin-left: 60px;
    padding: 20px;
    transition: margin-left 0.3s ease;
`;
