import React, { useEffect, useState } from 'react';
import axios from 'axios';
import ReactMarkdown from 'react-markdown';
import { MarkdownPreview } from '../../styles/CreateTaskStyles';
import { ContentContainer } from '../../styles/SidebarStyles';
import { Navigate, useParams } from 'react-router-dom';
import Sidebar from '../common/Sidebar';

const isAuthenticated = (): boolean => {
    return localStorage.getItem('authToken') !== null;
};

interface Task {
    title: string;
    deadline: string;
    priority: string;
    status: string;
    purpose: string;
    steps: string;
    memo: string;
    remarks: string;
}

const ViewTask: React.FC = () => {
    const { taskId } = useParams<{ taskId: string }>();
    const [task, setTask] = useState<Task | null>(null);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchTask = async () => {
            try {
                const response = await axios.get(`http://localhost:8000/api/tasks/${taskId}`);
                setTask(response.data);
            } catch (err) {
                setError('タスクの取得に失敗しました');
            }
        };

        fetchTask();
    }, [taskId]);

    if (!isAuthenticated()) {
        return <Navigate to="/Signin" replace />;
    }

    if (error) {
        return <p>{error}</p>;
    }

    return (
        <div>
            <Sidebar />
            <ContentContainer>
                <h2>タスク詳細ページ</h2>
                {task ? (
                    <MarkdownPreview>
                        <ReactMarkdown>{`
# タイトル
- ${task.title}

## 目的
${task.purpose}

## 期限
- ${task.deadline}

## 優先度
${task.priority}

## ステータス
${task.status}

### ステップ
${task.steps}

### メモ
${task.memo}

### 備考
${task.remarks}
                        `}</ReactMarkdown>
                    </MarkdownPreview>
                ) : (
                    <p>タスクを読み込んでいます...</p>
                )}
            </ContentContainer>
        </div>
    );
};

export default ViewTask;
