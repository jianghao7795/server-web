declare namespace ResetPassword {
  export type Password = {
    password: string;
    newPassword: string;
    repeatPassword: string;
  };

  export type Phone = {
    phone: string;
    newPhone: string;
    verification: string;
  };
}
