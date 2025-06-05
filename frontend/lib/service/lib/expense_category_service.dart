import 'dart:convert';
import 'package:http/http.dart' as http;
import '../base_service.dart' show BaseService;

class ExpenseCategoryService extends BaseService {
  Future<Map<String, dynamic>?> createExpenseCategory(
    Map<String, dynamic> data,
  ) async {
    try {
      final token = await getAuthToken();
      final response = await http.post(
        url('expense_categories/create'),
        headers: headersWithToken(token),
        body: jsonEncode(data),
      );

      final body = jsonDecode(response.body);
      return response.statusCode == 201
          ? body['expense_category']
          : {
            'error': body['message'] ?? 'Failed to create expense category',
            'statusCode': response.statusCode,
          };
    } catch (e) {
      return {'error': e.toString(), 'statusCode': 500};
    }
  }

  Future<List<Map<String, dynamic>>> getAllExpenseCategories() async {
    try {
      final token = await getAuthToken();
      final response = await http.get(
        url('expense_categories/'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return List<Map<String, dynamic>>.from(body['expense_categories']);
      } else {
        throw Exception('Failed to fetch expnense categories');
      }
    } catch (e) {
      return [];
    }
  }

  Future<Map<String, dynamic>?> getExpenseCategoryByID(int categoryID) async {
    try {
      final token = await getAuthToken();
      final response = await http.get(
        url('expense_categories/$categoryID'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return body['expense_category'];
      } else {
        return {
          'error': 'Failed to fetch expense category',
          'statusCode': response.statusCode,
        };
      }
    } catch (e) {
      return {'error': e.toString(), 'statusCode': 500};
    }
  }

  Future<List<Map<String, dynamic>>> getExpenseCategoryByUserID(
    int userID,
  ) async {
    try {
      final token = await getAuthToken();
      final response = await http.get(
        url('expense-categories/user/$userID'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return List<Map<String, dynamic>>.from(body['xpense_categories']);
      } else {
        throw Exception('Failed to fetch expense categores for user');
      }
    } catch (e) {
      return [];
    }
  }

  Future<Map<String, dynamic>?> updateExpenseCategory(
    int categoryID,
    Map<String, dynamic> data,
  ) async {
    try {
      final token = await getAuthToken();
      final response = await http.put(
        url('expense-categories/$categoryID'),
        headers: headersWithToken(token),
        body: jsonEncode(data),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return body['expense_category'];
      } else {
        return {
          'error': 'Failed to update expense category',
          'statusCode': response.statusCode,
        };
      }
    } catch (e) {
      return {'error': e.toString(), 'statusCode': 500};
    }
  }

  Future<bool> deleteExpenseCategory(int categoryID) async {
    try {
      final token = await getAuthToken();
      final response = await http.delete(
        url('expense-categories/$categoryID'),
        headers: headersWithToken(token),
      );

      return response.statusCode == 200;
    } catch (e) {
      return false;
    }
  }
}
