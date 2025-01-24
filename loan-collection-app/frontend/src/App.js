import React, { useState, useEffect } from 'react';
import { BrowserRouter, Routes, Route, Navigate, useNavigate } from 'react-router-dom';
import { ThemeProvider, createTheme, Box } from '@mui/material';
import LoginForm from './components/LoginForm';
import CasesList from './components/CasesList';
import CaseDetails from './components/CaseDetails';
import UserManagement from './components/UserManagement';
import SideNavigation from './components/SideNavigation';
import HomePage from './components/HomePage';
import { getUserRolesAndPermissions } from './service/api';
import { getAuthToken } from './service/auth';
import './App.css';
import AgencyManagement from './components/AgencyManagement';
import AgencyUserMapping from './components/AgencyUserMapping';
import CaseOnboarding from './components/CaseOnboarding';
import AgencyCaseManagement from './components/AgencyCaseManagement';
// Theme remains the same as in the previous App.js

const theme = createTheme({
  typography: {
    fontFamily: [
      'Inter', 
      '-apple-system', 
      'BlinkMacSystemFont', 
      '"Segoe UI"', 
      'Roboto', 
      '"Helvetica Neue"', 
      'Arial', 
      'sans-serif'
    ].join(','),
    h4: {
      fontWeight: 600,
    },
    h6: {
      fontWeight: 500,
    },
    body1: {
      fontWeight: 400,
    }
  },
  palette: {
    mode: 'dark',
    primary: {
      main: '#8a4fff',
    },
    background: {
      default: '#121212',
      paper: '#1e1e1e'
    }
  },
  components: {
    MuiButton: {
      styleOverrides: {
        root: {
          textTransform: 'none',
          borderRadius: 8,
        }
      }
    }
  }
});

const drawerWidth = 240;

function AuthenticatedContent({ onLogout, permissions }) {
  const navigate = useNavigate();

  useEffect(() => {
    navigate('/');
  }, []);

  return (
    <Box sx={{ display: 'flex' }}>
      <SideNavigation permissions={permissions} onLogout={onLogout} />
      <Box 
        component="main" 
        sx={{ 
          flexGrow: 1, 
          p: 3, 
          width: { sm: `calc(100% - ${drawerWidth}px)` } 
        }}
      >
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/cases" element={<CasesList />} />
          <Route path="/cases/:caseId" element={<CaseDetails />} />
          <Route path="/users" element={<UserManagement />} />
          <Route path="/agency-management" element={<AgencyManagement />} />
          <Route path="/agency-user-mapping" element={<AgencyUserMapping />} />
          <Route path="/case-onboarding" element={<CaseOnboarding />} />
          <Route path="/agency-cases" element={<AgencyCaseManagement />} />
          <Route path="*" element={<Navigate to="/" replace />} />
        </Routes>
      </Box>
    </Box>
  );
}

function App() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [permissions, setPermissions] = useState([]);

  const handleLogout = () => {
    localStorage.removeItem('token');
    setIsAuthenticated(false);
    setPermissions([]);
  };

  useEffect(() => {
    const checkAuth = async () => {
      const token = getAuthToken();
      if (token) {
        try {
          const userPermissions = await getUserRolesAndPermissions();
          setPermissions(userPermissions);
          setIsAuthenticated(true);
        } catch (error) {
          console.error('Authentication failed:', error);
          setIsAuthenticated(false);
        }
      }
    };
    checkAuth();
  }, []);

  return (
    <BrowserRouter>
      {isAuthenticated ? (
        <AuthenticatedContent 
          onLogout={handleLogout}
          permissions={permissions}
        />
      ) : (
        <LoginForm onLogin={() => setIsAuthenticated(true)} />
      )}
    </BrowserRouter>
  );
}

export default App;