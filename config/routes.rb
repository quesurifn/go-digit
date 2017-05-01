Rails.application.routes.draw do
  post 'login', to: 'plaid#exchange'
  get 'swagger' => redirect('/apidocs/api-docs.json') if Rails.env.development?
end
