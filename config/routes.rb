Rails.application.routes.draw do
  root 'dashboard#index'
  post ':exchange', to: 'plaid#exchange'
end
