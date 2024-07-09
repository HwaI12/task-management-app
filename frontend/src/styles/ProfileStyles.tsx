import styled from 'styled-components';

export const ContentContainer = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    margin-left: 60px;
    padding: 20px;
    transition: margin-left 0.3s ease;
`;

export const PrioritySection = styled.div`
    // margin-bottom: 40px;
    width: 50%;
    border-bottom: 1px solid #ccc;
    padding-bottom: 20px;
`;

export const StatusSection = styled.div`
    display: flex;
    justify-content: space-between;

    & > div {
        flex: 1;
        margin: 0 10px;
    }
`;

export const TaskCard = styled.div`
    border: 1px solid #ddd;
    padding: 10px;
    margin-bottom: 10px;
    border-radius: 4px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);

    h3 {
        margin-top: 0;
    }

    p {
        margin: 5px 0;
    }
`;

export const UserContainer = styled.div`
    // margin-bottom: 40px;
    display: flex;
    flex-direction: row;
    justify-content: center; // 中央揃え
    align-items: flex-start; // 上揃え
    // justify-content: flex-start; // 左揃えに変更
    // align-items: flex-end; // 下に揃える
    transition: margin-left 0.3s ease;
    border-bottom: 1px solid #ccc;
    padding: 50px 0 0 20px;
    width: 50%;
`;

export const TitleUserName = styled.p`
    font-size: 5rem;
    font-weight: bold;
    color: #333;
    text-align: left;
    margin: 0;
    line-height: 1.2; // 行の高さを調整
`;

export const TitleUserID = styled.p`
    font-size: 1.5rem;
    font-weight: bold;
    color: #333;
    text-align: left;
    margin: 0;
    align-self: flex-end;
    margin-bottom: 1rem;
    margin-left: 10px;
`;
