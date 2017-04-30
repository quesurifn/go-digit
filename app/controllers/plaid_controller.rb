class PlaidController < ApplicationController
  def exchange
    @token = params[:token]
    response = begin
      plaid_client.item.public_token.exchange(@token) 
    rescue Plaid::PlaidError => e
      render_error(e.error_message, Rack::Utils.status_code(:bad_request))
    end
    render json: response, status: Rack::Utils.status_code(:ok) 
  end
end
