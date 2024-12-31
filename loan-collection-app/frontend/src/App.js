import React, { useState, useEffect } from 'react';
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
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

function App() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [permissions, setPermissions] = useState([]);

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
    // <ThemeProvider theme={theme}>
      <BrowserRouter>
        {isAuthenticated ? (
          <Box sx={{ display: 'flex' }}>
            <SideNavigation permissions={permissions} />
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
                <Route path="/cases/:id" element={<CaseDetails />} />
                <Route path="/users" element={<UserManagement />} />
                <Route path="/agency-management" element={<AgencyManagement />} />
                <Route path="*" element={<Navigate to="/" replace />} />
              </Routes>
            </Box>
          </Box>
        ) : (
          <LoginForm onLogin={() => setIsAuthenticated(true)} />
        )}
      </BrowserRouter>
    // </ThemeProvider>
  );
}

export default App;