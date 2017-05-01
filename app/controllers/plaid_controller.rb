class PlaidController < ApplicationController
  respond_to :json

  def exchange
    @public_token = params[:token]
    response = begin
      plaid_client.item.public_token.exchange(@public_token) 
    rescue Plaid::PlaidError => e
      render_error(e.error_message, Rack::Utils.status_code(:bad_request))
    end
    # should probably actually store this to db
    session[:access_token] = response["access_token"]
    response = auth(session[:access_token])
    render json: response, status: Rack::Utils.status_code(:ok)
  end

  private
  def auth(access_token)
    render_error(StandardError.new("Invalid access token."), Rack::Utils.status_code(:unauthorized)) if access_token.empty?
    plaid_client.auth.get(access_token)
  end
end
