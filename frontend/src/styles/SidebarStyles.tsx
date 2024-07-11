import styled from 'styled-components';
import {
    IoHomeOutline,
    IoNotificationsOutline,
    IoPersonOutline,
    IoLogOutOutline,
    IoAddCircleOutline
} from "react-icons/io5";
import { Link } from 'react-router-dom';

export const SidebarContainer = styled.div`
    width: 70px;
    background-color: #f8f8f8;
    height: 100vh;
    box-shadow: 2px 0 5px rgba(0,0,0,0.1);
    display: flex;
    flex-direction: column;
    position: fixed;
    top: 0;
    left: 0;
    z-index: 1000;

    @media (max-width: 768px) {
        width: 100%;
        height: 70px;
        flex-direction: row;
        bottom: 0;
        top: auto;
    }
`;

export const SidebarHeader = styled.div`
    background-color: #fff;
    box-shadow: 0 2px 5px rgba(0,0,0,0.1);
    display: flex;
    justify-content: center;
    align-items: center;

    @media (max-width: 768px) {
        display: none;
    }
`;

export const SidebarMenu = styled.div`
    flex: 1;
    padding: 20px 0;
    overflow-y: auto;
    display: flex;
    flex-direction: column;
    align-items: center;

    @media (max-width: 768px) {
        flex-direction: row;
        padding: 0;
        overflow-y: hidden;
    }
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

    @media (max-width: 768px) {
        width: auto;
        margin-bottom: 0;
        flex: 1;
    }
`;

export const SidebarLogoImg = styled.img`
    width: 80px;
    height: 70px;
    object-fit: cover;

    @media (max-width: 768px) {
        display: none;
    }
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

export const StyledIoAddCircleOutline = styled(IoAddCircleOutline)`
    font-size: 24px;
    color: #333;
`;

export const ContentContainer = styled.div`
    margin-left: 60px;
    padding: 20px;
    transition: margin-left 0.3s ease;

    @media (max-width: 768px) {
        margin-left: 0;
        margin-bottom: 70px;
    }
`;

export const StyledLink = styled(Link)`
    color: inherit;
    text-decoration: none;
`;
