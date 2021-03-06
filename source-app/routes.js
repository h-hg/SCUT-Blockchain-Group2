var tuna = require('./controller.js');

module.exports = function(app) {
	app.get('/', function(req, res) {
		tuna.index(req, res);
	});
	app.get('/search', function(req, res) {
		tuna.search_html(req, res);
	});
	app.get('/business_search', function(req, res) {
		tuna.business_search(req, res);
	});
	app.get('/form', function(req, res) {
		tuna.form_(req, res);
	});
	app.get('/transit_form', function(req, res) {
		tuna.transit_form(req, res);
	});
	app.post('/re_form', function(req, res) {
		var function_name = 'addExpressInfo'
		tuna.re_form(req, res, function_name);
	});
	app.post('/re_transit_form', function(req, res) {
		var function_name = 'addTransferInfo'
		tuna.re_form(req, res, function_name);
	});
	app.get('/source/:id', function(req, res) {
		var function_name = 'getExpressInfo'
		tuna.get_tuna(req, res, function_name);
	});
	app.get('/transit/:id', function(req, res) {
		var function_name = 'getTransferInfo'
		tuna.get_tuna(req, res, function_name);
	});
}
