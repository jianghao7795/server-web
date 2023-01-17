/// <reference types="vite/client" />
declare interface Window {
  $message: MessageApiInjection;
  $notification: NotificationApiInjection;
  $loadingBar: LoadingBarApiInjection;
}
