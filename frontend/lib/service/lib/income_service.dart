import 'dart:convert';

import 'package:http/http.dart' as http;

import '../base_service.dart' show BaseService;

class IncomeService extends BaseService {
  Future<Map<String, dynamic>?> createIncome(Map<String, dynamic> data) async {
    try {
      final token = await getAuthToken();
      final response = await http.post(
        url('incomes/create'),
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

  Future<List<Map<String, dynamic>>> getAllIncomes() async {
    try {
      final token = await getAuthToken();
      final response = await http.get(
        url('incomes/'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return List<Map<String, dynamic>>.from(body['incomes']);
      } else {
        throw Exception('Failed to fetch incomes');
      }
    } catch (e) {
      return [];
    }
  }

  Future<Map<String, dynamic>?> getIncomeByID(int incomeID) async {
    try {
      final token = await getAuthToken();
      final response = await http.get(
        url('incomes/$incomeID'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return body['income'];
      } else {
        return {
          'error': 'Failed to fetch income',
          'statusCode': response.statusCode,
        };
      }
    } catch (e) {
      return {'error': e.toString(), 'statusCode': 500};
    }
  }

  Future<List<Map<String, dynamic>>> getIncomeByUserID(int userID) async {
    try {
      final token = await getAuthToken();
      final response = await http.get(
        url('incomes/user/$userID'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return List<Map<String, dynamic>>.from(body['incomes']);
      } else {
        throw Exception('Failed to fetch income for user');
      }
    } catch (e) {
      return [];
    }
  }

  Future<List<Map<String, dynamic>>> getIncomeBySourceID(int sourceID) async {
    try {
      final token = await getAuthToken();
      final response = await http.get(
        url('incomes/source/$sourceID'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return List<Map<String, dynamic>>.from(body['incomes']);
      } else {
        throw Exception('Failed to fetch income by source');
      }
    } catch (e) {
      return [];
    }
  }

  Future<Map<String, dynamic>?> updateIncomeSource(
    int incomeID,
    Map<String, dynamic> data,
  ) async {
    try {
      final token = await getAuthToken();
      final response = await http.put(
        url('incomes/$incomeID'),
        headers: headersWithToken(token),
        body: jsonEncode(data),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return body['income'];
      } else {
        return {
          'error': 'Failed to update income',
          'statusCode': response.statusCode,
        };
      }
    } catch (e) {
      return {'error': e.toString(), 'statusCode': 500};
    }
  }

  Future<bool> deleteIncomeSource(int incomeID) async {
    try {
      final token = await getAuthToken();
      final response = await http.delete(
        url('incomes/$incomeID'),
        headers: headersWithToken(token),
      );

      return response.statusCode == 200;
    } catch (e) {
      return false;
    }
  }
}
