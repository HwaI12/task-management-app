import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate, useParams, Navigate } from 'react-router-dom';
import Sidebar from '../common/Sidebar';
import { ContentContainer } from '../../styles/SidebarStyles';

// ユーザーが認証されているか確認する関数
const isAuthenticated = (): boolean => {
    return localStorage.getItem('authToken') !== null;
};

const CreateTask: React.FC = () => {
    const [errorMessage, setErrorMessage] = useState('');
    const [title, setTitle] = useState('');
    const [deadline, setDeadline] = useState('');
    const [priority, setPriority] = useState('高');
    const [status, setStatus] = useState('未着手');
    const [purpose, setPurpose] = useState('');
    const [description, setDescription] = useState('');
    const [steps, setSteps] = useState('');
    const [memo, setMemo] = useState('');
    const [remarks, setRemarks] = useState('');
    
    const { userId } = useParams<{ userId: string }>();

    const navigate = useNavigate();

    if (!isAuthenticated()) {
        return <Navigate to="/Signin" replace />;
    }

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setErrorMessage('');
        try {
            const response = await axios.post('http://localhost:8000/api/tasks', {
                user_id: userId,
                title,
                deadline,
                priority,
                status,
                purpose,
                description,
                steps,
                memo,
                remarks,
            });

            console.log('登録成功:', response.data);
            navigate('/profile');
        } catch (error) {
            if (axios.isAxiosError(error)) {
                console.error('登録エラー:', error.response?.data || error.message);
                setErrorMessage(error.response?.data?.message || '登録に失敗しました。');
            } else {
                console.error('予期しないエラー:', error);
                setErrorMessage('予期しないエラーが発生しました。');
            }
        }
    };

    return (
        <div>
            <Sidebar />
            <ContentContainer>
                <h2>タスク追加ページ</h2>
                <p> ユーザー {userId}!</p>
                <form onSubmit={handleSubmit}>
                    <div>
                        <label>タイトル</label>
                        <input
                            type="text"
                            value={title}
                            onChange={(e) => setTitle(e.target.value)}
                        />
                    </div>
                    <div>
                        <label>期限</label>
                        <input
                            type="date"
                            value={deadline}
                            onChange={(e) => setDeadline(e.target.value)}
                        />
                    </div>
                    <div>
                        <label>優先度</label>
                        <select value={priority} onChange={(e) => setPriority(e.target.value)}>
                            <option value="高">高</option>
                            <option value="中">中</option>
                            <option value="低">低</option>
                        </select>
                    </div>
                    <div>
                        <label>ステータス</label>
                        <select value={status} onChange={(e) => setStatus(e.target.value)}>
                            <option value="未着手">未着手</option>
                            <option value="進行中">進行中</option>
                            <option value="完了">完了</option>
                        </select>
                    </div>
                    <div>
                        <label>目的</label>
                        <textarea
                            value={purpose}
                            onChange={(e) => setPurpose(e.target.value)}
                        ></textarea>
                    </div>
                    <div>
                        <label>説明</label>
                        <textarea
                            value={description}
                            onChange={(e) => setDescription(e.target.value)}
                        ></textarea>
                    </div>
                    <div>
                        <label>ステップ</label>
                        <textarea
                            value={steps}
                            onChange={(e) => setSteps(e.target.value)}
                        ></textarea>
                    </div>
                    <div>
                        <label>メモ</label>
                        <textarea
                            value={memo}
                            onChange={(e) => setMemo(e.target.value)}
                        ></textarea>
                    </div>
                    <div>
                        <label>備考</label>
                        <textarea
                            value={remarks}
                            onChange={(e) => setRemarks(e.target.value)}
                        ></textarea>
                    </div>
                    {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
                    <button type="submit">タスク追加</button>
                </form>
            </ContentContainer>
        </div>
    );
};

export default CreateTask;
