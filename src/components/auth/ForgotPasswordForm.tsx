import React from 'react';
import { Form, Input, Button, Card, message } from 'antd';
import { MailOutlined } from '@ant-design/icons';
import { Link } from 'react-router-dom';

interface ForgotPasswordFormData {
  email: string;
}

const ForgotPasswordForm: React.FC = () => {
  const onFinish = (values: ForgotPasswordFormData) => {
    // TODO: Implement actual password reset logic
    console.log('Reset password requested for:', values.email);
    message.success(`Password reset link has been sent to ${values.email}!`);
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50">
      <Card className="w-full max-w-md">
        <h1 className="text-2xl font-bold text-center mb-6">Forgot Password</h1>
        <Form
          name="forgotPassword"
          onFinish={onFinish}
          layout="vertical"
        >
          <Form.Item
            name="email"
            rules={[
              { required: true, message: 'Please input your email!' },
              { type: 'email', message: 'Please enter a valid email!' }
            ]}
          >
            <Input prefix={<MailOutlined />} placeholder="Email" />
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit" className="w-full">
              Reset Password
            </Button>
          </Form.Item>

          <div className="text-center">
            <Link to="/login" className="text-blue-600">
              Back to Login
            </Link>
          </div>
        </Form>
      </Card>
    </div>
  );
};

export default ForgotPasswordForm;
