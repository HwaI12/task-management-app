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
    const [priority, setPriority] = useState('');
    const [status, setStatus] = useState('');
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

            console.log('Register successful:', response.data);
            navigate('/home');
        } catch (error) {
            if (axios.isAxiosError(error)) {
                console.error('Signup error:', error.response?.data || error.message);
                setErrorMessage(error.response?.data?.message || '登録に失敗しました。');
            } else {
                console.error('Unexpected error:', error);
                setErrorMessage('予期せぬエラーが発生しました。');
            }
        }
    };

    return (
        <div>
            <Sidebar />
            <ContentContainer>
                <h2>Add Task Page</h2>
                <p> User {userId}!</p>
                <form onSubmit={handleSubmit}>
                    <div>
                        <label>Title</label>
                        <input
                            type="text"
                            value={title}
                            onChange={(e) => setTitle(e.target.value)}
                        />
                    </div>
                    <div>
                        <label>Deadline</label>
                        <input
                            type="date"
                            value={deadline}
                            onChange={(e) => setDeadline(e.target.value)}
                        />
                    </div>
                    <div>
                        <label>Priority</label>
                        <select value={priority} onChange={(e) => setPriority(e.target.value)}>
                            <option value="high">High</option>
                            <option value="medium">Medium</option>
                            <option value="low">Low</option>
                        </select>
                    </div>
                    <div>
                        <label>Status</label>
                        <select value={status} onChange={(e) => setStatus(e.target.value)}>
                            <option value="not_started">Not started</option>
                            <option value="in_progress">In progress</option>
                            <option value="done">Done</option>
                        </select>
                    </div>
                    <div>
                        <label>Purpose</label>
                        <textarea
                            value={purpose}
                            onChange={(e) => setPurpose(e.target.value)}
                        ></textarea>
                    </div>
                    <div>
                        <label>Description</label>
                        <textarea
                            value={description}
                            onChange={(e) => setDescription(e.target.value)}
                        ></textarea>
                    </div>
                    <div>
                        <label>Steps</label>
                        <textarea
                            value={steps}
                            onChange={(e) => setSteps(e.target.value)}
                        ></textarea>
                    </div>
                    <div>
                        <label>Memo</label>
                        <textarea
                            value={memo}
                            onChange={(e) => setMemo(e.target.value)}
                        ></textarea>
                    </div>
                    <div>
                        <label>Remarks</label>
                        <textarea
                            value={remarks}
                            onChange={(e) => setRemarks(e.target.value)}
                        ></textarea>
                    </div>
                    {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
                    <button type="submit">Add Task</button>
                </form>
            </ContentContainer>
        </div>
    );
};

export default CreateTask;
