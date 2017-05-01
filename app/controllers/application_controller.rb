class ApplicationController < ActionController::API
  include ::ActionController::Cookies

  private

  def plaid_client
    PlaidClient.instance.plaid
  end

  def render_error(error, status)
    render json: error, status: status
  end
end
