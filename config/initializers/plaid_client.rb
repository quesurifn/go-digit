require 'plaid'
require 'singleton'

class PlaidClient
  include Singleton
  def plaid
    @client ||= 
      Plaid::Client.new(env: Rails.configuration.plaid.env,
                        client_id: Rails.configuration.plaid.client_id,
                        secret: Rails.configuration.plaid.secret,
                        public_key: Rails.configuration.plaid.public_key)
  end
end
