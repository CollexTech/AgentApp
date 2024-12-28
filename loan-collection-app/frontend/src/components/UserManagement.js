import React, { useState, useEffect, useCallback } from 'react';
import {
    Visibility,
    VisibilityOff,
    AddCircleOutline,
    MoreVert
} from '@mui/icons-material';
import {
    Container,
    Typography,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    Paper,
    Button,
    Dialog,
    DialogTitle,
    DialogContent,
    DialogActions,
    TextField,
    CircularProgress,
    Checkbox,
    FormControlLabel,
    InputAdornment,
    IconButton,
    Menu,
    MenuItem,
    FormGroup
} from '@mui/material';
import {
    fetchUsers,
    createUser,
    getAllAvailableRoles,
    assignRolesToUser
} from '../service/api';

const initialUserState = {
    username: '',
    password: ''
};

function UserManagement() {
    const [users, setUsers] = useState([]);
    const [availableRoles, setAvailableRoles] = useState([]);
    const [selectedUser, setSelectedUser] = useState(null);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState(null);
    const [showPassword, setShowPassword] = useState(false);
    const [isCreateUserDialogOpen, setIsCreateUserDialogOpen] = useState(false);
    const [newUser, setNewUser] = useState(initialUserState);

    // State for user actions menu
    const [anchorEl, setAnchorEl] = useState(null);
    const [selectedUserForAction, setSelectedUserForAction] = useState(null);
    const [isDeleteUserDialogOpen, setIsDeleteUserDialogOpen] = useState(false);
    const [roleSelectionOpen, setRoleSelectionOpen] = useState(false);
    const [selectedRoles, setSelectedRoles] = useState([]);

    const handleUserActionsMenuOpen = (event, user) => {
        event.stopPropagation();
        setAnchorEl(event.currentTarget);
        setSelectedUserForAction(user);
    };

    const handleUserActionsMenuClose = () => {
        setAnchorEl(null);
        setSelectedUserForAction(null);
        setRoleSelectionOpen(false);
    };

    const handleCloseDeleteUserDialog = () => {
        setIsDeleteUserDialogOpen(false);
        setSelectedUserForAction(null);
    };

    const handleDeleteUser = async () => {
        try {
            // TODO: Implement actual delete user API call
            console.log('Deleting user:', selectedUserForAction.username);
            await loadData();
            handleCloseDeleteUserDialog();
        } catch (error) {
            console.error('Error deleting user:', error);
        }
    };

    const handleCloseUserDetails = () => {
        setSelectedUser(null);
    };

    const togglePasswordVisibility = () => {
        setShowPassword(prev => !prev);
    };

    const loadData = useCallback(async (isMounted) => {
        try {
            setIsLoading(true);
            const [fetchedUsers, fetchedRoles] = await Promise.all([
                fetchUsers(),
                getAllAvailableRoles()
            ]);

            if (isMounted) {
                setUsers(fetchedUsers);
                setAvailableRoles(fetchedRoles);
            }
        } catch (error) {
            console.error('Error fetching data:', error);
            if (isMounted) {
                setError(error);
            }
        } finally {
            if (isMounted) {
                setIsLoading(false);
            }
        }
    }, []);

    useEffect(() => {
        let isMounted = true;
        loadData(isMounted);
        return () => {
            isMounted = false;
        };
    }, [loadData]);

    const handleCreateUser = async () => {
        try {
            await createUser({
                username: newUser.username,
                password: newUser.password
            });
            await loadData(true);
            handleCreateUserClose();
        } catch (error) {
            console.error('User creation failed:', error);
        }
    };

    const handleCreateUserOpen = () => setIsCreateUserDialogOpen(true);
    const handleCreateUserClose = () => {
        setIsCreateUserDialogOpen(false);
        setNewUser(initialUserState);
    };

    const handleAssignRoles = async (userId, selectedRoles) => {
        try {
            await assignRolesToUser(userId, selectedRoles);
            await loadData(true);
        } catch (error) {
            console.error('Failed to assign roles:', error);
        }
    };

    const handleRoleChange = (roleId) => {
        setSelectedRoles(prev => {
            if (prev.includes(roleId)) {
                return prev.filter(id => id !== roleId);
            }
            return [...prev, roleId];
        });
    };

    if (isLoading) return <CircularProgress />;
    if (error) return <Typography color="error">Error loading data</Typography>;

    return (
        <Container>
            <Typography variant="h4">User Management</Typography>
            <Button
                startIcon={<AddCircleOutline />}
                onClick={handleCreateUserOpen}
            >
                Create User
            </Button>

            {/* User Creation Dialog */}
            <Dialog open={isCreateUserDialogOpen} onClose={handleCreateUserClose}>
                <DialogTitle>Create New User</DialogTitle>
                <DialogContent>
                    <TextField
                        name="username"
                        label="Username"
                        value={newUser.username}
                        onChange={(e) => setNewUser(prev => ({...prev, username: e.target.value}))}
                        fullWidth
                        margin="normal"
                    />
                    <TextField
                        name="password"
                        label="Password"
                        type={showPassword ? 'text' : 'password'}
                        value={newUser.password}
                        onChange={(e) => setNewUser(prev => ({...prev, password: e.target.value}))}
                        fullWidth
                        margin="normal"
                        InputProps={{
                            endAdornment: (
                                <InputAdornment position="end">
                                    <IconButton onClick={togglePasswordVisibility} edge="end">
                                        {showPassword ? <VisibilityOff /> : <Visibility />}
                                    </IconButton>
                                </InputAdornment>
                            )
                        }}
                    />
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleCreateUserClose}>Cancel</Button>
                    <Button onClick={handleCreateUser}>Create</Button>
                </DialogActions>
            </Dialog>

            {/* User Actions Menu */}
            <Menu
                anchorEl={anchorEl}
                open={Boolean(anchorEl)}
                onClose={handleUserActionsMenuClose}
                anchorOrigin={{
                    vertical: 'bottom',
                    horizontal: 'right',
                }}
                transformOrigin={{
                    vertical: 'top',
                    horizontal: 'right',
                }}
            >
                <MenuItem 
                    onClick={() => {
                        setRoleSelectionOpen(true);
                        setAnchorEl(null);
                    }}
                >
                    Assign Roles
                </MenuItem>
                <MenuItem 
                    onClick={() => {
                        setIsDeleteUserDialogOpen(true);
                        setAnchorEl(null);
                    }}
                >
                    Delete User
                </MenuItem>
            </Menu>

            {/* Role Selection Dialog - Move outside of Menu */}
            <Dialog 
                open={roleSelectionOpen} 
                onClose={() => {
                    setRoleSelectionOpen(false);
                    setSelectedRoles([]);
                }}
            >
                <DialogTitle>Assign Roles</DialogTitle>
                <DialogContent>
                    <FormGroup>
                        {availableRoles.map((role) => (
                            <FormControlLabel
                                key={role.ID}
                                control={
                                    <Checkbox
                                        checked={selectedUserForAction?.role_list.includes(role.RoleName) || selectedRoles.includes(role.RoleName)}
                                        disabled={selectedUserForAction?.role_list.includes(role.RoleName)}
                                        onChange={() => handleRoleChange(role.RoleName)}
                                    />
                                }
                                label={role.RoleName}
                            />
                        ))}
                    </FormGroup>
                </DialogContent>
                <DialogActions>
                    <Button onClick={() => {
                        setRoleSelectionOpen(false);
                        setSelectedRoles([]);
                    }}>
                        Cancel
                    </Button>
                    <Button onClick={() => {
                        handleAssignRoles(selectedUserForAction.id, selectedRoles);
                        setRoleSelectionOpen(false);
                        setSelectedRoles([]);
                    }}>
                        Save
                    </Button>
                </DialogActions>
            </Dialog>

            {/* Delete User Confirmation Dialog */}
            <Dialog
                open={isDeleteUserDialogOpen}
                onClose={handleCloseDeleteUserDialog}
                maxWidth="xs"
                fullWidth
            >
                <DialogTitle>Delete User</DialogTitle>
                <DialogContent>
                    <Typography>
                        Are you sure you want to delete user {selectedUserForAction?.username}?
                        This action cannot be undone.
                    </Typography>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleCloseDeleteUserDialog} color="secondary">
                        Cancel
                    </Button>
                    <Button
                        onClick={handleDeleteUser}
                        color="error"
                        variant="contained"
                    >
                        Delete
                    </Button>
                </DialogActions>
            </Dialog>

            {selectedUser && (
                <Dialog open={!!selectedUser} onClose={handleCloseUserDetails}>
                    <DialogTitle>User Details</DialogTitle>
                    <DialogContent>
                        <Typography>Username: {selectedUser.username}</Typography>
                        <Typography>Email: {selectedUser.email || 'N/A'}</Typography>
                        <Typography>Roles: {(selectedUser.role_list || []).join(', ')}</Typography>
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={handleCloseUserDetails}>Close</Button>
                    </DialogActions>
                </Dialog>
            )}

            {/* User List Table */}
            <TableContainer component={Paper}>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell>Username</TableCell>
                            <TableCell>Email</TableCell>
                            <TableCell>Roles</TableCell>
                            <TableCell>Actions</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {users.map(user => (
                            <TableRow key={user.id}>
                                <TableCell>{user.username}</TableCell>
                                <TableCell>{user.email}</TableCell>
                                <TableCell>{(user.role_list || []).join(', ')}</TableCell>
                                <TableCell>
                                    <IconButton
                                        aria-label="user actions"
                                        onClick={(e) => handleUserActionsMenuOpen(e, user)}
                                    >
                                        <MoreVert />
                                    </IconButton>
                                </TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </Container>
    );
}

export default UserManagement;