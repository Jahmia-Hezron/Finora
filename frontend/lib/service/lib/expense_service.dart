import 'dart:convert';

import 'package:http/http.dart' as http;

import '../base_service.dart' show BaseService;

class ExpenseService extends BaseService {
  Future<Map<String, dynamic>?> createExpense(Map<String, dynamic> data) async {
    try {
      final token = await getAuthToken();
      final response = await http.post(
        url('expenses/create'),
        headers: headersWithToken(token),
        body: jsonEncode(data),
      );

      final body = jsonDecode(response.body);
      return response.statusCode == 200
          ? body
          : {
            'error': body['message'] ?? 'Failed to create expense',
            'statusCode': response.statusCode,
          };
    } on Exception catch (e) {
      return {'error': e.toString(), 'statusCode': 500};
    }
  }

  Future<List<Map<String, dynamic>>> getAllExpenses() async {
    try {
      final token = await getAuthToken();
      final response = await http.get(
        url('expenses/'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return List<Map<String, dynamic>>.from(body['expenses']);
      } else {
        throw Exception('Failed to fetch expenses');
      }
    } catch (e) {
      return [];
    }
  }

  Future<Map<String, dynamic>?> getExpenseByID(int expnenseID) async {
    try {
      final token = await getAuthToken();
      final response = await http.get(
        url('expenses/$expnenseID'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return body['expense'];
      } else {
        return {
          'error': 'Failed to fetch expense',
          'statusCode': response.statusCode,
        };
      }
    } catch (e) {
      return {'error': e.toString(), 'statusCode': 500};
    }
  }

  Future<List<Map<String, dynamic>>> getExpensesByUserID(int userID) async {
    try {
      final token = await getAuthToken();
      final response = await http.get(
        url('expenses/user/$userID'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return List<Map<String, dynamic>>.from(body['expenses']);
      } else {
        throw Exception('Failed to fetch expenses for user');
      }
    } catch (e) {
      return [];
    }
  }

  Future<List<Map<String, dynamic>>> getExpenseByCategoryID(
    int categoryID,
  ) async {
    try {
      final token = await getAuthToken();
      final response = await http.get(
        url('expenses/category/$categoryID'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return List<Map<String, dynamic>>.from(body['expenses']);
      } else {
        throw Exception('Failed to fetch expenses by source');
      }
    } catch (e) {
      return [];
    }
  }

  Future<Map<String, dynamic>?> updateExpense(
    int expenseID,
    Map<String, dynamic> data,
  ) async {
    try {
      final token = await getAuthToken();
      final response = await http.put(
        url('expenses/$expenseID'),
        headers: headersWithToken(token),
        body: jsonEncode(data),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return body['expense'];
      } else {
        return {
          'error': 'Failed to update expense',
          'statusCode': response.statusCode,
        };
      }
    } catch (e) {
      return {'error': e.toString(), 'statusCode': 500};
    }
  }

  Future<bool> deleteExpense(int expenseID) async {
    try {
      final token = await getAuthToken();
      final response = await http.delete(
        url('expenses/$expenseID'),
        headers: headersWithToken(token),
      );

      return response.statusCode == 200;
    } catch (e) {
      return false;
    }
  }
}
