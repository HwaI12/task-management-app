import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useParams, Link } from 'react-router-dom';
import Sidebar from '../common/Sidebar';
import { ContentContainer, PrioritySection, StatusSection, TaskCard, UserContainer, TitleUserName, TitleUserID } from '../../styles/ProfileStyles';

// Taskの型定義
interface Task {
    id: number; // IDフィールドは数値であることが一般的
    title: string;
    purpose: string;
    deadline: string;
    priority: 'high' | 'middle' | 'low';
    status: 'yet' | 'progress' | 'done';
}

// Userの型定義
interface User {
    user_id: string;
    username: string;
}

// 認証されたユーザーIDを取得する関数
const getAuthenticatedUserId = (): string | null => {
    return localStorage.getItem('userId');
};

const Profile: React.FC = () => {
    const { userId } = useParams<{ userId: string }>();
    const [tasks, setTasks] = useState<Task[] | null>(null);
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState(true);
    const authenticatedUserId = getAuthenticatedUserId();

    useEffect(() => {
        const fetchTasks = async () => {
            try {
                const response = await axios.get(`http://localhost:8000/api/tasks?user_id=${userId}`);
                setTasks(response.data);
            } catch (error) {
                console.error('タスクの取得に失敗しました:', error);
            }
        };

        const fetchUser = async () => {
            try {
                const response = await axios.get(`http://localhost:8000/api/user?user_id=${userId}`);
                setUser(response.data);
            } catch (error) {
                console.error('ユーザー情報の取得に失敗しました:', error);
            } finally {
                setLoading(false);
            }
        };

        fetchTasks();
        fetchUser();
    }, [userId]);

    if (loading) {
        return <div>Loading...</div>;
    }

    if (!user) {
        return (
            <ContentContainer>
                <Sidebar />
                このアカウントは存在しません
            </ContentContainer>
        );
    }

    // タスクを優先度とステータスごとに分類して表示
    const renderTasks = (priority: string, status: string) => {
        return tasks
            ? tasks
                .filter(task => task.priority === priority && task.status === status)
                .map(task => (
                    <TaskCard key={task.id}>
                        <Link to={`/${userId}/task/${task.id}`}>
                            <h3>{task.title}</h3>
                            {/* -----を表示 */}
                            <hr />
                            <p>{task.deadline}まで</p>
                        </Link>
                    </TaskCard>
                ))
            : null;
    };

    return (
        <div>
            <Sidebar />
            <ContentContainer>
                <UserContainer>
                    <TitleUserName>{user.username}</TitleUserName>
                    <TitleUserID>@{userId}</TitleUserID>
                </UserContainer>

                {userId === authenticatedUserId ? (
                    <>
                        <PrioritySection>
                            <h3>高優先度</h3>
                            <StatusSection>
                                <div>
                                    <h4>未着手</h4>
                                    {renderTasks('high', 'yet')}
                                </div>
                                <div>
                                    <h4>進行中</h4>
                                    {renderTasks('high', 'progress')}
                                </div>
                                <div>
                                    <h4>完了</h4>
                                    {renderTasks('high', 'done')}
                                </div>
                            </StatusSection>
                        </PrioritySection>

                        <PrioritySection>
                            <h3>中優先度</h3>
                            <StatusSection>
                                <div>
                                    <h4>未着手</h4>
                                    {renderTasks('middle', 'yet')}
                                </div>
                                <div>
                                    <h4>進行中</h4>
                                    {renderTasks('middle', 'progress')}
                                </div>
                                <div>
                                    <h4>完了</h4>
                                    {renderTasks('middle', 'done')}
                                </div>
                            </StatusSection>
                        </PrioritySection>

                        <PrioritySection>
                            <h3>低優先度</h3>
                            <StatusSection>
                                <div>
                                    <h4>未着手</h4>
                                    {renderTasks('low', 'yet')}
                                </div>
                                <div>
                                    <h4>進行中</h4>
                                    {renderTasks('低', '進行中')}
                                </div>
                                <div>
                                    <h4>完了</h4>
                                    {renderTasks('low', 'done')}
                                </div>
                            </StatusSection>
                        </PrioritySection>
                    </>
                ) : (
                    <div>
                        <h3>制作物</h3>
                        <p>ここに認証されたユーザー以外のコンテンツが表示されます。</p>
                    </div>
                )}
            </ContentContainer>
        </div>
    );
};

export default Profile;
