import React, { useState, useEffect } from 'react';
import { BrowserRouter, Routes, Route, Navigate, Link } from 'react-router-dom';
import { ThemeProvider, createTheme } from '@mui/material/styles';
import { 
  Box, 
  Drawer, 
  List, 
  ListItem, 
  ListItemText, 
  Typography, 
  Toolbar, 
  ListItemIcon,
  Container,
  Paper
} from '@mui/material';
import { 
  Home as HomeIcon, 
  ListAlt as CaseIcon, 
  PermIdentity as PermissionIcon 
} from '@mui/icons-material';
import LoginForm from './components/LoginForm';
import CasesList from './components/CasesList';
import CaseDetails from './components/CaseDetails';
import { getUserRolesAndPermissions } from './service/api';
import { getAuthToken } from './service/auth';
import './App.css';

// Create a custom Material-UI theme
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

function HomePage() {
  return (
    <Container maxWidth="md" sx={{ mt: 4 }}>
      <Paper elevation={3} sx={{ p: 4, textAlign: 'center' }}>
        <Typography variant="h4" gutterBottom>
          Welcome to Loan Collection App
        </Typography>
        <Typography variant="body1" paragraph>
          This is your central dashboard for managing loan collection cases.
        </Typography>
      </Paper>
    </Container>
  );
}

function SideNavigation({ permissions }) {
  console.log('Permissions received:', permissions); // Debug log

  return (
    <Drawer
      variant="permanent"
      sx={{
        width: drawerWidth,
        flexShrink: 0,
        [`& .MuiDrawer-paper`]: { width: drawerWidth, boxSizing: 'border-box' },
      }}
    >
      <Toolbar />
      <Box sx={{ overflow: 'auto' }}>
        {/* Permissions Section */}
        <Typography variant="h6" sx={{ p: 2 }}>
          My Permissions
        </Typography>
        <List>
          {permissions && permissions.length > 0 ? (
            permissions.map((permission, index) => (
              <ListItem key={index}>
                <ListItemIcon>
                  <PermissionIcon />
                </ListItemIcon>
                <ListItemText 
                  primary={permission} 
                  primaryTypographyProps={{ 
                    variant: 'body2',
                    style: { 
                      textTransform: 'capitalize',
                      fontWeight: 400 
                    } 
                  }}
                />
              </ListItem>
            ))
          ) : (
            <ListItem>
              <ListItemText 
                primary="No permissions found" 
                primaryTypographyProps={{ 
                  variant: 'body2',
                  color: 'textSecondary' 
                }}
              />
            </ListItem>
          )}
        </List>

        {/* Navigation Section */}
        <Typography variant="h6" sx={{ p: 2 }}>
          Navigation
        </Typography>
        <List>
          <ListItem component={Link} to="/home" button>
            <ListItemIcon>
              <HomeIcon />
            </ListItemIcon>
            <ListItemText primary="Home" />
          </ListItem>
          <ListItem component={Link} to="/cases" button>
            <ListItemIcon>
              <CaseIcon />
            </ListItemIcon>
            <ListItemText primary="My Cases" />
          </ListItem>
        </List>
      </Box>
    </Drawer>
  );
}

function App() {
  const [permissions, setPermissions] = useState([]);
  const [isAuthenticated, setIsAuthenticated] = useState(!!getAuthToken());

  useEffect(() => {
    const fetchPermissions = async () => {
      try {
        if (isAuthenticated) {
          const data = await getUserRolesAndPermissions();
          console.log('Fetched permissions data:', data); // Debug log
          
          // Ensure we're setting the correct permissions
          const fetchedPermissions = Array.isArray(data) 
            ? data 
            : (data.permissions || []);
          
          setPermissions(fetchedPermissions);
        }
      } catch (error) {
        console.error('Failed to fetch permissions', error);
        setPermissions([]);
      }
    };

    fetchPermissions();
  }, [isAuthenticated]);

  return (
    <ThemeProvider theme={theme}>
      <BrowserRouter>
        <Box sx={{ display: 'flex' }}>
          {isAuthenticated && <SideNavigation permissions={permissions} />}
          <Box 
            component="main" 
            sx={{ 
              flexGrow: 1, 
              p: 3, 
              width: { sm: `calc(100% - ${drawerWidth}px)` },
              ml: isAuthenticated ? `${drawerWidth}px` : 0 
            }}
          >
            <Routes>
              <Route 
                path="/" 
                element={isAuthenticated ? <Navigate to="/home" /> : <LoginForm />} 
              />
              <Route 
                path="/home" 
                element={isAuthenticated ? <HomePage /> : <Navigate to="/" />} 
              />
              <Route 
                path="/cases" 
                element={isAuthenticated ? <CasesList /> : <Navigate to="/" />} 
              />
              <Route 
                path="/cases/:id" 
                element={isAuthenticated ? <CaseDetails /> : <Navigate to="/" />} 
              />
            </Routes>
          </Box>
        </Box>
      </BrowserRouter>
    </ThemeProvider>
  );
}

export default App;