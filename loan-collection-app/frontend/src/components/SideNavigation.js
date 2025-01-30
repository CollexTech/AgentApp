import React, { useMemo } from 'react';
import { 
  Drawer, 
  List, 
  ListItem, 
  ListItemText, 
  ListItemIcon,
  Box,
  Typography
} from '@mui/material';
import { 
  Home as HomeIcon, 
  ListAlt as CaseIcon, 
  People as UsersIcon,
  Business as BuildingIcon,
  GroupAdd as GroupAddIcon,
  Upload as UploadFile,
  Logout as LogoutIcon
} from '@mui/icons-material';
import { Link } from 'react-router-dom';

const drawerWidth = 240;

const listItems = {
  "home": {
    label: 'Home',
    url: "/",
    icon: <HomeIcon />
  },
  "view_my_cases": {
    label: 'My Cases',
    url: "/cases",
    icon: <CaseIcon />
  },
  "view_users": {
    label: 'User Management',
    url: "/users",
    icon: <UsersIcon />
  },
  "view_agencies": {
    label: 'Agency Management',
    url: "/agency-management",
    icon: <BuildingIcon />
  },
  "view_agency_user_mapping": {
    label: 'Agency User Mapping',
    url: "/agency-user-mapping",
    icon: <GroupAddIcon />
  },
  "assign_agency_cases": {
    label: 'Case Onboarding',
    url: "/case-onboarding",
    icon: <UploadFile />
  },
  "view_agency_cases": {
    label: 'Agency Cases',
    url: "/agency-cases",
    icon: <BuildingIcon />
  },
}

function SideNavigation({ permissions, onLogout }) {
  const sideNavigationItems = useMemo(() => (
    permissions.reduce((acc, permission) => {
      const item = listItems[permission];
      if (item) acc.push(item);
      return acc;
    }, [])
  ), [permissions])

  return (
    <Drawer
      variant="permanent"
      sx={{
        width: drawerWidth,
        flexShrink: 0,
        [`& .MuiDrawer-paper`]: { width: drawerWidth, boxSizing: 'border-box' },
      }}
    >
      <Box sx={{ 
        overflow: 'auto',
        display: 'flex',
        flexDirection: 'column',
        height: '100%'
      }}>
        {/* Navigation Section */}
        <Typography variant="h6" sx={{ p: 2 }}>
          Navigation
        </Typography>
        <List>
          {sideNavigationItems.map((item, index) => (
            <ListItem button component={Link} to={item.url} key={index}>
              <ListItemIcon>
                {item.icon}
              </ListItemIcon>
              <ListItemText primary={item.label} />
            </ListItem>
          ))}
        </List>
        <Box sx={{ marginTop: 'auto' }}>
          <List>
            <ListItem onClick={onLogout}>
              <ListItemIcon>
                <LogoutIcon />
              </ListItemIcon>
              <ListItemText primary="Logout" />
            </ListItem>
          </List>
        </Box>
      </Box>
    </Drawer>
  );
}

export default SideNavigation;