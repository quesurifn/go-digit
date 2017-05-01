Rails.application.routes.draw do
  post ':exchange', to: 'plaid#exchange'
end
