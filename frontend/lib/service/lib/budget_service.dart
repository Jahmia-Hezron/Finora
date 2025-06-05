import 'dart:convert';

import 'package:http/http.dart' as http;

import '../base_service.dart' show BaseService;

class BudgetService extends BaseService {
  Future<Map<String, dynamic>?> createBudget(Map<String, dynamic> data) async {
    try {
      final token = await getAuthToken();
      final response = await http.post(
        url('budgets/create'),
        headers: headersWithToken(token),
        body: jsonEncode(data),
      );

      final body = jsonDecode(response.body);
      return response.statusCode == 200
          ? body
          : {
            'error': body['message'] ?? 'Failed to create income',
            'statusCode': response.statusCode,
          };
    } on Exception catch (e) {
      return {'error': e.toString(), 'statusCode': 500};
    }
  }

  Future<List<Map<String, dynamic>>> getAllBudgets() async {
    try {
      final token = await getAuthToken();
      final response = await http.get(
        url('budgets/'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return List<Map<String, dynamic>>.from(body['budgets']);
      } else {
        throw Exception('Failed to fetch budgets');
      }
    } catch (e) {
      return [];
    }
  }

  Future<Map<String, dynamic>?> getBudgetByID(int budgetID) async {
    try {
      final token = await getAuthToken();
      final response = await http.get(
        url('budgets/$budgetID'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return body['budget'];
      } else {
        return {
          'error': 'Failed to fetch budget',
          'statusCode': response.statusCode,
        };
      }
    } catch (e) {
      return {'error': e.toString(), 'statusCode': 500};
    }
  }

  Future<List<Map<String, dynamic>>> getBudgetByUserID(int userID) async {
    try {
      final token = await getAuthToken();
      final response = await http.get(
        url('budgets/user/$userID'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return List<Map<String, dynamic>>.from(body['budgets']);
      } else {
        throw Exception('Failed to fetch budgets for user');
      }
    } catch (e) {
      return [];
    }
  }

  Future<Map<String, dynamic>?> getBudgetsByCategoryID(int categoryID) async {
    try {
      final token = await getAuthToken();
      final response = await http.get(
        url('budgets/$categoryID'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return body['budgets'];
      } else {
        return {
          'error': 'Failed to fetch budgets by category',
          'statusCode': response.statusCode,
        };
      }
    } catch (e) {
      return {'error': e.toString(), 'statusCode': 500};
    }
  }

  Future<Map<String, dynamic>?> updateBudget(
    int incomeID,
    Map<String, dynamic> data,
  ) async {
    try {
      final token = await getAuthToken();
      final response = await http.put(
        url('budgets/$incomeID'),
        headers: headersWithToken(token),
        body: jsonEncode(data),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return body['budget'];
      } else {
        return {
          'error': 'Failed to update budget',
          'statusCode': response.statusCode,
        };
      }
    } catch (e) {
      return {'error': e.toString(), 'statusCode': 500};
    }
  }

  Future<bool> deleteBudget(int budgetID) async {
    try {
      final token = await getAuthToken();
      final response = await http.delete(
        url('budgets/$budgetID'),
        headers: headersWithToken(token),
      );

      return response.statusCode == 200;
    } catch (e) {
      return false;
    }
  }
}
