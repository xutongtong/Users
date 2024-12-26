import React from 'react';
import { Form, Input, Button, Card, message } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';
import { Link, useNavigate } from 'react-router-dom';

interface LoginFormData {
  username: string;
  password: string;
}

const LoginForm: React.FC = () => {
  const navigate = useNavigate();

  const onFinish = (values: LoginFormData) => {
    // TODO: Implement actual login logic
    if (values.username === 'admin' && values.password === 'admin') {
      message.success('Login successful!');
      localStorage.setItem('isAuthenticated', 'true');
      navigate('/dashboard');
    } else {
      message.error('Invalid credentials!');
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50">
      <Card className="w-full max-w-md">
        <h1 className="text-2xl font-bold text-center mb-6">Admin Login</h1>
        <Form
          name="login"
          initialValues={{ remember: true }}
          onFinish={onFinish}
          layout="vertical"
        >
          <Form.Item
            name="username"
            rules={[{ required: true, message: 'Please input your username!' }]}
          >
            <Input prefix={<UserOutlined />} placeholder="Username" />
          </Form.Item>

          <Form.Item
            name="password"
            rules={[{ required: true, message: 'Please input your password!' }]}
          >
            <Input.Password prefix={<LockOutlined />} placeholder="Password" />
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit" className="w-full">
              Log in
            </Button>
          </Form.Item>

          <div className="text-center">
            <Link to="/forgot-password" className="text-blue-600">
              Forgot password?
            </Link>
            <span className="mx-2">|</span>
            <Link to="/register" className="text-blue-600">
              Register now!
            </Link>
          </div>
        </Form>
      </Card>
    </div>
  );
};

export default LoginForm;
