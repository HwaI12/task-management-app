// Title
// SubTitle
// SubTitle
// SubTitle
// みたいな感じで入力する

// タスクをスムーズに行うためにはどのようなSubTitleが必要か考えて
// Title: タスク名
// SubTitle:期限(日付)
// SubTitle:優先度(高・中・低)
// SubTitle:ステータス(未着手、進行中、完)
// SubTitle:タスクの目的
// SubTitle:タスクの説明
// SubTitle:ステップ(タスクを達成するためのステップ)
// SubTitle:メモ
// SubTitle:備考

import React from 'react';
import { useParams } from 'react-router-dom';
import Sidebar from '../common/Sidebar';
import { ContentContainer } from '../../styles/SidebarStyles';

const AddTask: React.FC = () => {
    const { userId } = useParams<{ userId: string }>();

    return (
        <div>
            <ContentContainer>
            <Sidebar />
            <h2>AddTask Page</h2>
            <p> User {userId}!</p>
            </ContentContainer>

        </div>
    );
};

export default AddTask;
