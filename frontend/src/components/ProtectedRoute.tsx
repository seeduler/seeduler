import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import api from '../api';

interface ProtectedRouteProps {
    children: React.ReactNode;
}

export const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ children }) => {
    const [hasHalls, setHasHalls] = useState<boolean | null>(null);
    const navigate = useNavigate();

    useEffect(() => {
        const checkHalls = async () => {
            try {
                const response = await api.get('/halls');
                setHasHalls(true);
            } catch (error) {
                if (error.response?.status === 412) { // PreconditionFailed
                    setHasHalls(false);
                    navigate('/upload-data');
                }
            }
        };
        checkHalls();
    }, [navigate]);

    if (hasHalls === null) {
        return <div>Loading...</div>;
    }

    if (hasHalls === false) {
        return <div>Please upload initial data first</div>;
    }

    return <>{children}</>;
}; 