import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { ConfigProvider, Layout, Menu } from 'antd';
import LoginForm from './components/auth/LoginForm';
const { Sider, Content } = Layout;
import RegisterForm from './components/auth/RegisterForm';
import ForgotPasswordForm from './components/auth/ForgotPasswordForm';
import UserManagement from './components/users/UserManagement';

const ProtectedRoute: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const isAuthenticated = localStorage.getItem('isAuthenticated') === 'true';
  return isAuthenticated ? <>{children}</> : <Navigate to="/login" />;
};

const App: React.FC = () => {
  return (
    <ConfigProvider>
      <Layout className="min-h-screen">
        <Router>
          <Routes>
            <Route path="/login" element={<LoginForm />} />
            <Route path="/register" element={<RegisterForm />} />
            <Route path="/forgot-password" element={<ForgotPasswordForm />} />
            <Route
              path="/dashboard"
              element={
                <ProtectedRoute>
                  <Layout className="flex h-screen">
                    <Sider className="bg-white border-r border-gray-200 shadow-sm">
                      <div className="p-4 border-b border-gray-200">
                        <h1 className="text-xl font-semibold text-gray-800">Admin Dashboard</h1>
                      </div>
                      <Menu mode="inline" defaultSelectedKeys={['dashboard']} className="mt-2">
                        <Menu.Item key="dashboard" className="text-gray-700">Dashboard</Menu.Item>
                        <Menu.Item key="users" className="text-gray-700">User Management</Menu.Item>
                      </Menu>
                    </Sider>
                    <Layout className="bg-gray-50">
                      <Content className="p-6 m-4 bg-white rounded-lg shadow-sm">
                        <UserManagement />
                      </Content>
                    </Layout>
                  </Layout>
                </ProtectedRoute>
              }
            />
            <Route path="/" element={<Navigate to="/login" />} />
          </Routes>
        </Router>
      </Layout>
    </ConfigProvider>
  );
};

export default App;
